import type { Project } from "@/generated/types/models/project";

export interface ProjectTableProps {
  data: Project[];
  onRowClick?: (item: Project) => void;
}

export function ProjectTable({ data, onRowClick }: ProjectTableProps) {
  return (
    <table className="w-full border-collapse" data-testid="projects-table">
      <thead>
        <tr className="border-b">
          <th className="px-4 py-2 text-left text-sm font-medium">Code</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Description</th>
          <th className="px-4 py-2 text-left text-sm font-medium">ID</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Name</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Organization ID</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item) => (
          <tr
            key={item.id}
            data-testid="project-row"
            onClick={() => onRowClick?.(item)}
            className="border-b hover:bg-muted/50 cursor-pointer"
          >
            <td className="px-4 py-2 text-sm">{item.code}</td>
            <td className="px-4 py-2 text-sm">{item.description}</td>
            <td className="px-4 py-2 text-sm">{item.id}</td>
            <td className="px-4 py-2 text-sm">{item.name}</td>
            <td className="px-4 py-2 text-sm">{item.organizationID}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
