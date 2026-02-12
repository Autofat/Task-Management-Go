import { useState } from "react";
import Input from "../atoms/Input";
import Button from "../atoms/Button";

const ProjectForm = ({ onSubmit, isLoading, initialData = null }) => {
  const [formData, setFormData] = useState(
    initialData || {
      title: "",
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

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Input
        type="text"
        label="Project Title"
        name="title"
        value={formData.title}
        onChange={handleChange}
        placeholder="Enter project title"
        required
      />
      <Button type="submit" variant="primary" disabled={isLoading}>
        {isLoading
          ? "Saving..."
          : initialData
            ? "Update Project"
            : "Create Project"}
      </Button>
    </form>
  );
};

export default ProjectForm;
