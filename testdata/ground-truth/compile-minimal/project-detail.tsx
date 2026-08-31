import type { Project } from "@/generated/types/models/project";

export interface ProjectDetailProps {
  data: Project;
}

export function ProjectDetail({ data }: ProjectDetailProps) {
  return (
    <dl className="space-y-4" data-testid="project-detail">
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Code</dt>
        <dd className="text-sm" data-testid="project-code">{data.code}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Description</dt>
        <dd className="text-sm" data-testid="project-description">{data.description}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">ID</dt>
        <dd className="text-sm" data-testid="project-id">{data.id}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Name</dt>
        <dd className="text-sm" data-testid="project-name">{data.name}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Organization ID</dt>
        <dd className="text-sm" data-testid="project-organization-id">{data.organizationID}</dd>
      </div>
    </dl>
  );
}
