import { mockAnomalies } from '../mocks/anomalies';

export async function getAnomalies(projectId: string) {
  return mockAnomalies.filter(a => a.project_id === projectId);
}
