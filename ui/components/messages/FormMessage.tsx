"use client";

import { useState } from "react";
import type { FormMessage as FormMessageType, FormField } from "@/lib/types";

interface FormMessageProps {
  message: FormMessageType;
  onSubmit: (id: string, data: Record<string, unknown>) => void;
}

export function FormMessage({ message, onSubmit }: FormMessageProps) {
  const [values, setValues] = useState<Record<string, unknown>>(() => {
    const defaults: Record<string, unknown> = {};
    for (const field of message.fields) {
      if (field.default !== undefined) {
        defaults[field.name] = field.default;
      } else if (field.type === "boolean") {
        defaults[field.name] = false;
      } else {
        defaults[field.name] = "";
      }
    }
    return defaults;
  });
  const [submitted, setSubmitted] = useState(false);

  function handleChange(name: string, value: unknown) {
    setValues((prev) => ({ ...prev, [name]: value }));
  }

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setSubmitted(true);
    onSubmit(message.id, values);
  }

  if (submitted) {
    return (
      <div className="text-[13px] italic flex items-center gap-2" style={{ color: "#545f7e" }}>
        <div className="w-[7px] h-[7px] rounded-full" style={{ backgroundColor: "#2D864E" }} />
        Form submitted. Waiting for response...
      </div>
    );
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-5">
      <div style={{ borderBottom: "1px solid #DEE1EA", paddingBottom: "12px" }}>
        <h3 className="text-[14px] font-semibold" style={{ color: "#111626" }}>
          {message.title}
        </h3>
        {message.description && (
          <p className="text-[12px] mt-1 leading-relaxed" style={{ color: "#545f7e" }}>
            {message.description}
          </p>
        )}
      </div>

      <div className="space-y-4" style={{ maxWidth: "428px" }}>
        {message.fields.map((field) => (
          <FieldInput
            key={field.name}
            field={field}
            value={values[field.name]}
            onChange={(v) => handleChange(field.name, v)}
          />
        ))}
      </div>

      <button
        type="submit"
        className="px-5 py-2 rounded text-[13px] font-medium text-white transition-colors"
        style={{ backgroundColor: "#1F7A78" }}
        onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#135757")}
        onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "#1F7A78")}
      >
        {message.submit_label}
      </button>
    </form>
  );
}

function FieldInput({
  field,
  value,
  onChange,
}: {
  field: FormField;
  value: unknown;
  onChange: (v: unknown) => void;
}) {
  const inputStyle: React.CSSProperties = {
    width: "100%",
    border: "1px solid #E1DDDA",
    borderRadius: "4px",
    backgroundColor: "#ffffff",
    padding: "8px 12px",
    fontSize: "13px",
    color: "#111626",
    minHeight: "40px",
    outline: "none",
  };

  return (
    <div>
      <label
        className="block text-[12px] font-medium mb-1.5"
        style={{ color: "#111626", letterSpacing: "0.02em" }}
      >
        {field.label}
        {field.required && <span style={{ color: "#DD314F", marginLeft: "2px" }}>*</span>}
      </label>
      {field.description && (
        <p className="text-[11px] mb-1.5 leading-relaxed" style={{ color: "#545f7e" }}>
          {field.description}
        </p>
      )}

      {field.type === "select" || field.enum ? (
        <select
          value={String(value ?? "")}
          onChange={(e) => onChange(e.target.value)}
          required={field.required}
          style={inputStyle}
        >
          <option value="">Select...</option>
          {(field.enum ?? []).map((opt) => (
            <option key={opt} value={opt}>{opt}</option>
          ))}
        </select>
      ) : field.type === "boolean" ? (
        <label className="flex items-center gap-2.5 text-[13px] cursor-pointer" style={{ color: "#111626" }}>
          <input
            type="checkbox"
            checked={Boolean(value)}
            onChange={(e) => onChange(e.target.checked)}
            className="w-4 h-4 rounded"
            style={{ borderColor: "#E1DDDA" }}
          />
          {field.label}
        </label>
      ) : field.type === "textarea" ? (
        <textarea
          value={String(value ?? "")}
          onChange={(e) => onChange(e.target.value)}
          required={field.required}
          placeholder={field.placeholder}
          rows={3}
          style={inputStyle}
        />
      ) : (
        <input
          type={field.type === "password" ? "password" : field.type === "number" ? "number" : "text"}
          value={String(value ?? "")}
          onChange={(e) => onChange(field.type === "number" ? Number(e.target.value) : e.target.value)}
          required={field.required}
          placeholder={field.placeholder}
          style={inputStyle}
        />
      )}
    </div>
  );
}
