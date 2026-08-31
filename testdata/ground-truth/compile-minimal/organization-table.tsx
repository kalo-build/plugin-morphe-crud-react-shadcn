import type { Organization } from "@/generated/types/models/organization";

export interface OrganizationTableProps {
  data: Organization[];
  onRowClick?: (item: Organization) => void;
}

export function OrganizationTable({ data, onRowClick }: OrganizationTableProps) {
  return (
    <table className="w-full border-collapse" data-testid="organizations-table">
      <thead>
        <tr className="border-b">
          <th className="px-4 py-2 text-left text-sm font-medium">Code</th>
          <th className="px-4 py-2 text-left text-sm font-medium">ID</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Name</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item) => (
          <tr
            key={item.id}
            data-testid="organization-row"
            onClick={() => onRowClick?.(item)}
            className="border-b hover:bg-muted/50 cursor-pointer"
          >
            <td className="px-4 py-2 text-sm">{item.code}</td>
            <td className="px-4 py-2 text-sm">{item.id}</td>
            <td className="px-4 py-2 text-sm">{item.name}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
