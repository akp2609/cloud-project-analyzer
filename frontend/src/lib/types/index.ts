export type Environment = 'DEV' | 'STAGING' | 'PROD';

export type Project = {
  id: string;
  name: string;
  environment: Environment;

  features: {
    cost_enabled: boolean;
    logs_enabled: boolean;
    metrics_enabled: boolean;
    slo_enabled: boolean;
  };
};

export type CostAnomaly = {
  id: string;
  project_id: string;
  service: string;
  date: string;

  actual_cost: number;
  baseline_cost: number;
  spike_percentage: number;

  severity: 'LOW' | 'MEDIUM' | 'HIGH';
};

export type ProjectInsight = {
  id: string;
  project_id: string;

  title: string;
  description: string;

  severity: 'LOW' | 'MEDIUM' | 'HIGH';
  created_at: string;

  related_anomaly_id?: string;
};
