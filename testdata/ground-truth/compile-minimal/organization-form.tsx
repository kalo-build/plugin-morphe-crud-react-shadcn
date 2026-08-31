"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { Organization } from "@/generated/types/models/organization";
import { OrganizationSchema } from "@/generated/schemas/models/organization";

export interface OrganizationFormProps {
  defaultValues?: Partial<Organization>;
  onSubmit: (data: Organization) => void | Promise<void>;
  disabled?: boolean;
}

export function OrganizationForm({ defaultValues, onSubmit, disabled }: OrganizationFormProps) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Organization>({
    resolver: zodResolver(OrganizationSchema),
    defaultValues,
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4" data-testid="organization-form">
      <div>
        <label htmlFor="code" className="block text-sm font-medium">
          Code
        </label>
        <input
          id="code"
          type="text"
          data-testid="organization-code-input"
          {...register("code")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.code && (
          <p className="mt-1 text-sm text-destructive">{errors.code.message}</p>
        )}
      </div>
      <div>
        <label htmlFor="name" className="block text-sm font-medium">
          Name
        </label>
        <input
          id="name"
          type="text"
          data-testid="organization-name-input"
          {...register("name")}
          disabled={disabled || isSubmitting}
          className="mt-1 block w-full rounded-md border border-input px-3 py-2"
        />
        {errors.name && (
          <p className="mt-1 text-sm text-destructive">{errors.name.message}</p>
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
