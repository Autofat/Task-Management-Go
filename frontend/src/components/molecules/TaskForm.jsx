import { useState } from "react";
import Input from "../atoms/Input";
import Textarea from "../atoms/Textarea";
import Select from "../atoms/Select";
import Button from "../atoms/Button";

const TaskForm = ({
  onSubmit,
  isLoading,
  members = [],
  initialData = null,
}) => {
  const [formData, setFormData] = useState(
    initialData || {
      title: "",
      description: "",
      priority: "medium",
      assigned_id: "",
      due_date: "",
    },
  );

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  const priorityOptions = [
    { value: "low", label: "Low" },
    { value: "medium", label: "Medium" },
    { value: "high", label: "High" },
  ];

  const memberOptions = members.map((member) => ({
    value: member.user_id,
    label: member.fullname,
  }));

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Input
        type="text"
        label="Task Title"
        name="title"
        value={formData.title}
        onChange={handleChange}
        placeholder="Enter task title"
        required
      />
      <Textarea
        label="Description"
        name="description"
        value={formData.description}
        onChange={handleChange}
        placeholder="Enter task description"
      />
      <Select
        label="Priority"
        name="priority"
        value={formData.priority}
        onChange={handleChange}
        options={priorityOptions}
        required
      />
      <Select
        label="Assign To"
        name="assigned_id"
        value={formData.assigned_id}
        onChange={handleChange}
        options={memberOptions}
        required
      />
      <Input
        type="date"
        label="Due Date"
        name="due_date"
        value={formData.due_date}
        onChange={handleChange}
      />
      <Button type="submit" variant="primary" disabled={isLoading}>
        {isLoading ? "Saving..." : initialData ? "Update Task" : "Create Task"}
      </Button>
    </form>
  );
};

export default TaskForm;
