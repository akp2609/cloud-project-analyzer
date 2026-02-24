import { ProjectInsight } from '../types';

export const mockInsights: ProjectInsight[] = [
  {
    id: 'ins-1',
    project_id: 'proj-1',
    title: 'BigQuery cost spike detected',
    description:
      'BigQuery usage exceeded baseline by 166%. Consider reviewing scheduled queries.',
    severity: 'HIGH',
    created_at: '2026-01-28T10:00:00Z',
    related_anomaly_id: 'anom-1',
  },
];
