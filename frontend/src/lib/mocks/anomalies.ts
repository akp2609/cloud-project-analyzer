import { CostAnomaly } from '../types';

export const mockAnomalies: CostAnomaly[] = [
  {
    id: 'anom-1',
    project_id: 'proj-1',
    service: 'BigQuery',
    date: '2026-01-28',
    actual_cost: 120,
    baseline_cost: 45,
    spike_percentage: 166,
    severity: 'HIGH',
  },
];
