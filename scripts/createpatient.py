import json
import uuid
import numpy as np
import requests

names = ['chin', 'jia', 'xiong', 'juin', 'tian', 'ng', 'kam', 'woh', 'desmond', 'yeoh', 'khooi', 'xin', 'zhe',
         'chan', 'chee', 'seng']
data = []

"""
{
	"name": "Xiao Ming",
	"phoneNumber": "01223287423",
    "id": "951221081234",
    "status": 1,
    "email": "xiaoming@um.edu.my"
}
"""

url = 'https://api.staging.cosmos.care:443/v1/patients'

for i in range(200):
    telegram_id = str(uuid.uuid4())
    name = ' '.join([name.capitalize() for name in np.random.choice(names, 3, replace=False)])
    phonenumber = ''.join([str(np.random.randint(0, 10)) for _ in range(10)])
    id = ''.join([str(np.random.randint(10, 100)) for _ in range(12)])
    status = np.random.randint(1, 6)
    email = ''.join(name.split(' ')).lower() + '@gg.com'

    body = {
        'name': name,
        'phoneNumber': phonenumber,
        'id': id,
        'status': status,
        'email': email
    }
    data = {'data': body}

    jsonstr = json.dumps(data)
    headers = {'Authorization': 'g380egN18Gew8WXryqTAEi4YQyh2'}
    print(requests.post(url + '/' + id, data=jsonstr, headers=headers))
   