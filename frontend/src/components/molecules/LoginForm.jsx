import { useState } from "react";
import Input from "../atoms/Input";
import Button from "../atoms/Button";

const LoginForm = ({ onSubmit, isLoading }) => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
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
        placeholder="Enter your password"
        required
      />
      <Button type="submit" variant="primary" disabled={isLoading}>
        {isLoading ? "Loading..." : "Login"}
      </Button>
    </form>
  );
};

export default LoginForm;
