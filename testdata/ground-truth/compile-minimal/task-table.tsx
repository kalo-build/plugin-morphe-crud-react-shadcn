import type { Task } from "@/generated/types/models/task";

export interface TaskTableProps {
  data: Task[];
  onRowClick?: (item: Task) => void;
}

export function TaskTable({ data, onRowClick }: TaskTableProps) {
  return (
    <table className="w-full border-collapse" data-testid="tasks-table">
      <thead>
        <tr className="border-b">
          <th className="px-4 py-2 text-left text-sm font-medium">ID</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Status</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Title</th>
          <th className="px-4 py-2 text-left text-sm font-medium">Project ID</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item) => (
          <tr
            key={item.id}
            data-testid="task-row"
            onClick={() => onRowClick?.(item)}
            className="border-b hover:bg-muted/50 cursor-pointer"
          >
            <td className="px-4 py-2 text-sm">{item.id}</td>
            <td className="px-4 py-2 text-sm">{item.status}</td>
            <td className="px-4 py-2 text-sm">{item.title}</td>
            <td className="px-4 py-2 text-sm">{item.projectID}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
