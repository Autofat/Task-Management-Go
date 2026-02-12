import { useState } from "react";
import Input from "../atoms/Input";
import Button from "../atoms/Button";

const RegisterForm = ({ onSubmit, isLoading }) => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    fullname: "",
  });

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
    <form onSubmit={handleSubmit} className="space-y-4 w-full max-w-md">
      <Input
        type="text"
        label="Full Name"
        name="fullname"
        value={formData.fullname}
        onChange={handleChange}
        placeholder="Enter your full name"
        required
      />
      <Input
        type="email"
        label="Email"
        name="email"
        value={formData.email}
        onChange={handleChange}
        placeholder="Enter your email"
        required
      />
      <Input
        type="password"
        label="Password"
        name="password"
        value={formData.password}
        onChange={handleChange}
        placeholder="Enter your password (min 6 characters)"
        required
      />
      <Button type="submit" variant="primary" disabled={isLoading}>
        {isLoading ? "Loading..." : "Register"}
      </Button>
    </form>
  );
};

export default RegisterForm;
