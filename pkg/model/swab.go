package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/utility"
	"context"
	"fmt"
	"github.com/twinj/uuid"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSwab creates new swab
func (m *Model) CreateSwab(ctx context.Context, swab *dto.Swab, patientType int64, user *dto.User) (*dto.Swab, error) {

	// some processing for swab
	swab.ID = utility.GenerateSwabID(swab.PatientID, swab.Date)

	// check if swab exist
	_, err := m.swabDAO.Get(ctx, swab.ID, constants.AllPatients)

	// only can create swab if not found
	if err != nil && status.Code(err) == codes.Unknown {

		// check if patient exist, put patient data
		p, err := m.GetPatient(ctx, swab.PatientID, patientType)
		if err != nil {
			return nil, constants.PatientNotFoundError
		}
		swab.PatientName = p.Name
		swab.PatientPhoneNumber = p.PhoneNumber
		swab.PatientType = p.Type

		// create swab
		s, err := m.swabDAO.Create(ctx, swab)
		if err != nil {
			return nil, err
		}

		// update patient
		p.SwabCount = p.SwabCount + 1
		if swab.Date > p.SwabDate || p.SwabDate == "" {
			p.SwabDate = swab.Date
			// calculate days since swab if needed
			if p.SwabDate == "" {
				p.DaysSinceSwab = 0
			} else {
				t, err := utility.DateStringToTime(p.SwabDate)
				if err != nil {
					return nil, err
				}
				p.DaysSinceSwab = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
			}
		}
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("failed to update patient , id: %s | patientId: %s", swab.ID, swab.PatientID))
		}

		// add into activity if user not nil
		if user != nil {
			_, err = m.createSwabActivity(ctx, nil, swab, user)
			if err != nil {
				return nil, err
			}
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, status.Error(codes.AlreadyExists, "Swab already exist!")
}

// UpdateSwab updates swab
func (m *Model) UpdateSwab(ctx context.Context, swab *dto.Swab, patientType int64, user *dto.User) (*dto.Swab, error) {
	// get swabDate
	patientID, _ := utility.GetPatientIDAndDateFromSwabID(swab.ID)

	// check if swab exist
	s, err := m.swabDAO.Get(ctx, swab.ID, patientType)
	if err != nil {
		return nil, err
	}

	// tech debt: get old swab
	oldSwab, err := m.swabDAO.Get(ctx, swab.ID, patientType)
	if err != nil {
		return nil, err
	}

	// patch swab
	s.Status = swab.Status
	s.Location = swab.Location
	s.IsOtherSwabLocation = swab.IsOtherSwabLocation

	// check if patient exist
	_, err = m.GetPatient(ctx, patientID, constants.AllPatients)
	if err != nil {
		return nil, constants.PatientNotFoundError
	}

	// update swab
	_, err = m.swabDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	// add into activity if user not nil
	if user != nil {
		_, err = m.createSwabActivity(ctx, oldSwab, s, user)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

// UpdateSwabs update swabs
func (m *Model) UpdateSwabs(ctx context.Context, swab *dto.Swab, ids []string, patientType int64, user *dto.User) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	swab.ID = ids[0]
	s, err := m.UpdateSwab(ctx, swab, patientType, user)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetSwab gets swab by ID
func (m *Model) GetSwab(ctx context.Context, id string, patientType int64) (*dto.Swab, error) {
	return m.swabDAO.Get(ctx, id, patientType)
}

// BatchGetSwabs get swabs by slice of IDs
func (m *Model) BatchGetSwabs(ctx context.Context, ids []string, patientType int64) ([]*dto.Swab, error) {
	return m.swabDAO.BatchGet(ctx, ids, patientType)
}

// QuerySwabs queries swabs by sort, range, filter
func (m *Model) QuerySwabs(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Swab, error) {
	return m.swabDAO.Query(ctx, sort, itemsRange, filter, patientType)
}

// DeleteSwab deletes swab by ID
func (m *Model) DeleteSwab(ctx context.Context, id string, patientType int64) (*dto.Swab, error) {
	// check if swab exist
	s, err := m.swabDAO.Get(ctx, id, patientType)
	if err != nil {
		return nil, err
	}

	// check if patient exist
	patientID, _ := utility.GetPatientIDAndDateFromSwabID(id)
	p, err := m.GetPatient(ctx, patientID, constants.AllPatients)
	if err != nil {
		return nil, constants.PatientNotFoundError
	}

	// delete swab
	err = m.swabDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	// update patient
	p.SwabCount = p.SwabCount - 1
	if s.Date == p.SwabDate {
		p.SwabDate = ""
		p.DaysSinceSwab = 0
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("failed to update patient , id: %s | patientId: %s", id, patientID))
		}
	}
	return s, nil
}

// DeleteSwabs delete swabs by IDs
func (m *Model) DeleteSwabs(ctx context.Context, ids []string, patientType int64) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteSwab(ctx, id, patientType)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}

// QuerySwabsByPatientID ...
func (m *Model) QuerySwabsByPatientID(ctx context.Context, id string, patientType int64) ([]*dto.Swab, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	_, swabs, err := m.swabDAO.Query(ctx, nil, nil, filter, patientType)
	return swabs, err
}

// createSwabActivity creates new swab activity
func (m *Model) createSwabActivity(ctx context.Context, oldSwab, newSwab *dto.Swab, user *dto.User) (*dto.Activity, error) {
	activity := &dto.Activity{
		ID:       uuid.NewV4().String(),
		UserID:   user.ID,
		UserName: user.DisplayName,
		OldSwab:  oldSwab,
		NewSwab:  newSwab,
		Time:     utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		TTL:      utility.MilliToTime(time.Now().Add(time.Hour*24*constants.ActivityTTLDays).Unix()*1000 - 1000),
	}

	// create activity
	s, err := m.activityDAO.Create(ctx, activity)
	if err != nil {
		return nil, err
	}

	// return result
	return s, nil
}
