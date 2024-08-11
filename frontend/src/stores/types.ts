export interface Task {
  id: string;
  title: string;
  description: string;
  checked: boolean;
  priority: string;
  children: Task[];
}

export enum Priority {
  DEFAULT = "default",
  LOW = "low",
  MEDIUM = "medium",
  HIGH = "high"
} 