import type { Organization } from "@/generated/types/models/organization";

export interface OrganizationDetailProps {
  data: Organization;
}

export function OrganizationDetail({ data }: OrganizationDetailProps) {
  return (
    <dl className="space-y-4" data-testid="organization-detail">
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Code</dt>
        <dd className="text-sm" data-testid="organization-code">{data.code}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">ID</dt>
        <dd className="text-sm" data-testid="organization-id">{data.id}</dd>
      </div>
      <div>
        <dt className="text-sm font-medium text-muted-foreground">Name</dt>
        <dd className="text-sm" data-testid="organization-name">{data.name}</dd>
      </div>
    </dl>
  );
}
