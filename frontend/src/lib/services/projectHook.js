import { GoogleAuth } from 'google-auth-library';

const BASE_URL = process.env.PROJECT_HOOK_URL;

async function getClient() {
  const auth = new GoogleAuth();
  return await auth.getIdTokenClient(BASE_URL);
}

export async function linkProject(raw){
    const client = await getClient();

    const res = await client.request({
        url: `${BASE_URL}/projects/link`,
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        data: raw,
    })
    return res.data;
}