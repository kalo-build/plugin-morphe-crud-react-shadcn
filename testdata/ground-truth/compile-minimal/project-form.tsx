"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { Project } from "@/generated/types/models/project";
import { ProjectSchema } from "@/generated/schemas/models/project";

export interface ProjectFormProps {
  defaultValues?: Partial<Project>;
  onSubmit: (data: Project) => void | Promise<void>;
  disabled?: boolean;
}

export function ProjectForm({ defaultValues, onSubmit, disabled }: ProjectFormProps) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Project>({
    resolver: zodResolver(ProjectSchema),
    defaultValues,
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4" data-testid="project-form">
      <div>
        <label htmlFor="code" className="block text-sm font-medium">
          Code
        </label>
        <input
          id="code"
          type="text"
          data-testid="project-code-input"
          {...register("code")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.code && (
          <p className="mt-1 text-sm text-destructive">{errors.code.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="description" className="block text-sm font-medium">
          Description
        </label>
        <input
          id="description"
          type="text"
          data-testid="project-description-input"
          {...register("description")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.description && (
          <p className="mt-1 text-sm text-destructive">{errors.description.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="name" className="block text-sm font-medium">
          Name
        </label>
        <input
          id="name"
          type="text"
          data-testid="project-name-input"
          {...register("name")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.name && (
          <p className="mt-1 text-sm text-destructive">{errors.name.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="organizationID" className="block text-sm font-medium">
          Organization ID
        </label>
        <input
          id="organizationID"
          type="text"
          data-testid="project-organization-id-input"
          {...register("organizationID")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.organizationID && (
          <p className="mt-1 text-sm text-destructive">{errors.organizationID.message}</p>
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
