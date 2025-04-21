import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://13.233.49.27:3000/user';

  const randomId = Math.floor(Math.random() * 1000) + 1;

  const payload = JSON.stringify({
    id: randomId,
    name: 'John Doe',
    email: 'john@example.com',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  console.log('Status code:', res.status);
  console.log('Response body:', res.body);
  console.log('Response time was', res.timings.duration + 'ms');

  check(res, {
    'is status 201': (r) => r.status === 201,
  });
}
