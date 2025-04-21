import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 10,
  duration: '30s',
};

export default function () {
  const payload = JSON.stringify({
    id: Math.floor(Math.random() * 10000), // generate a random ID-
    name: 'John Doe',
    email: 'john@example.com'
  });

  const headers = { 'Content-Type': 'application/json' };

  let res = http.post('http://13.233.49.27:3000/user', payload, { headers });

  check(res, {
    'is status 201': (r) => r.status === 201,
  });

  console.log(`Response body: ${res.body}`);
  sleep(1);
}
