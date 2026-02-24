import { GoogleAuth } from 'google-auth-library';

const BASE_URL = process.env.ANALYSIS_ENGINE_URL;

async function getClient() {
  const auth = new GoogleAuth();
  return await auth.getIdTokenClient(BASE_URL);
}


export async function fetchAnalysisEngine(path, params = {}) {
  const client = await getClient();

  const query = new URLSearchParams(params).toString();
  const url = `${BASE_URL}${path}${query ? `?${query}` : ''}`;

  const res = await client.request({ url });
  return res.data;
}


export async function fetchDashboard(projectId) {
  const client = await getClient();
  const res = await client.request({
    url: `${BASE_URL}/dashboard?project_id=${projectId}`,
  });
  return res.data;
}


export async function fetchMetricsInsight(projectId, metricType = "") {
  const params = { project_id: projectId };
  if (metricType) {
    params.metric_type = metricType;
  }
  return await fetchAnalysisEngine("/projects/metrics-insight", params);
}


export async function fetchAnomalies(projectId) {
  return await fetchAnalysisEngine("/projects/anomalies", { project_id: projectId });
}


export async function fetchInsights(projectId) {
  return await fetchAnalysisEngine("/projects/insights", { project_id: projectId });
}
