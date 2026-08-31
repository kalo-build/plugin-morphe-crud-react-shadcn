import type { Task } from "@/generated/types/models/task";

export interface TaskDetailProps {
  data: Task;
}

export function TaskDetail({ data }: TaskDetailProps) {
  return (
    <dl className="space-y-4" data-testid="task-detail">
      <div>
        <dt className="text-sm font-medium text-muted-foreground">ID</dt>
        <dd className="text-sm" data-testid="task-id">{data.id}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Status</dt>
        <dd className="text-sm" data-testid="task-status">{data.status}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Title</dt>
        <dd className="text-sm" data-testid="task-title">{data.title}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Project ID</dt>
        <dd className="text-sm" data-testid="task-project-id">{data.projectID}</dd>
      </div>
    </dl>
  );
}
