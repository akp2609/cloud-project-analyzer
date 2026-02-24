import { GoogleAuth } from 'google-auth-library';
import fs from 'fs';  
import FormData from 'form-data'; 

const BASE_URL = process.env.UPLOAD_SERVICE_URL;

async function getClient() {
  const auth = new GoogleAuth();
  return await auth.getIdTokenClient(BASE_URL);
}


export async function uploadFile(file, filename, headers) {
  const client = await getClient();

  const form = new FormData();
  form.append('file', file, filename);

  const res = await client.request({
    url: `${BASE_URL}/upload`, 
    method: 'POST',
    headers: {
      ...headers,
      ...form.getHeaders(), 
    },
    data: form,
  });

  return res.data;
}
