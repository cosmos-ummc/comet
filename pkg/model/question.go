package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateQuestion creates new question
func (m *Model) CreateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error) {

	// check if question exist
	_, err := m.questionDAO.Get(ctx, question.ID)

	// only can create question if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create question
		s, err := m.questionDAO.Create(ctx, question)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.QuestionAlreadyExistError
}

// UpdateQuestion updates question
func (m *Model) UpdateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error) {
	// check if question exist
	s, err := m.questionDAO.Get(ctx, question.ID)
	if err != nil {
		return nil, err
	}

	// patch question
	s.Content = question.Content
	s.Type = question.Type
	s.Category = question.Category

	// update question
	_, err = m.questionDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateQuestions update questions
func (m *Model) UpdateQuestions(ctx context.Context, question *dto.Question, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	question.ID = ids[0]
	s, err := m.UpdateQuestion(ctx, question)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetQuestion gets question by ID
func (m *Model) GetQuestion(ctx context.Context, id string) (*dto.Question, error) {
	return m.questionDAO.Get(ctx, id)
}

// BatchGetQuestions get questions by slice of IDs
func (m *Model) BatchGetQuestions(ctx context.Context, ids []string) ([]*dto.Question, error) {
	return m.questionDAO.BatchGet(ctx, ids)
}

// QueryQuestions queries questions by sort, range, filter
func (m *Model) QueryQuestions(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Question, error) {
	return m.questionDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteQuestion deletes question by ID
func (m *Model) DeleteQuestion(ctx context.Context, id string) (*dto.Question, error) {
	// check if question exist
	s, err := m.questionDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete question
	err = m.questionDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteQuestions delete questions by IDs
func (m *Model) DeleteQuestions(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteQuestion(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}

// QueryQuestionsByPatientID ...
func (m *Model) QueryQuestionsByPatientID(ctx context.Context, id string) ([]*dto.Question, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	_, questions, err := m.questionDAO.Query(ctx, nil, nil, filter)
	return questions, err
}
