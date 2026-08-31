"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { Task } from "@/generated/types/models/task";
import { TaskSchema } from "@/generated/schemas/models/task";

export interface TaskFormProps {
  defaultValues?: Partial<Task>;
  onSubmit: (data: Task) => void | Promise<void>;
  disabled?: boolean;
}

export function TaskForm({ defaultValues, onSubmit, disabled }: TaskFormProps) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Task>({
    resolver: zodResolver(TaskSchema),
    defaultValues,
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4" data-testid="task-form">
      <div>
        <label htmlFor="status" className="block text-sm font-medium">
          Status
        </label>
        <input
          id="status"
          type="text"
          data-testid="task-status-input"
          {...register("status")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.status && (
          <p className="mt-1 text-sm text-destructive">{errors.status.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="title" className="block text-sm font-medium">
          Title
        </label>
        <input
          id="title"
          type="text"
          data-testid="task-title-input"
          {...register("title")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.title && (
          <p className="mt-1 text-sm text-destructive">{errors.title.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="projectID" className="block text-sm font-medium">
          Project ID
        </label>
        <input
          id="projectID"
          type="text"
          data-testid="task-project-id-input"
          {...register("projectID")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.projectID && (
          <p className="mt-1 text-sm text-destructive">{errors.projectID.message}</p>
        )}
      </div>
      <button
        type="submit"
        data-testid="submit-button"
        disabled={disabled || isSubmitting}
        className="rounded-md bg-primary px-4 py-2 text-primary-foreground"
      >
        {isSubmitting ? "Saving..." : "Save"}
      </button>
    </form>
  );
}
