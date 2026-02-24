import { mockInsights } from '../mocks/insights';

export async function getInsights(projectId: string) {
  return mockInsights.filter(i => i.project_id === projectId);
}
