import { Project } from '../types';

export const mockProjects: Project[] = [
  {
    id: 'proj-1',
    name: 'Billing Analyzer',
    environment: 'DEV',
    features: {
      cost_enabled: true,
      logs_enabled: true,
      metrics_enabled: true,
      slo_enabled: false,
    },
  },
];
