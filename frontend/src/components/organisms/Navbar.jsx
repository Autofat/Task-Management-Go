import { useNavigate, Link } from "react-router-dom";
import Button from "../atoms/Button";

const Navbar = () => {
  const navigate = useNavigate();
  const token = localStorage.getItem("token");
  const user = JSON.parse(localStorage.getItem("user") || "{}");

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    navigate("/login");
  };

  return (
    <nav className="bg-white shadow-md">
      <div className="container mx-auto px-4 py-4">
        <div className="flex justify-between items-center">
          <Link to="/" className="text-2xl font-bold text-blue-600">
            Task Management
          </Link>

          {token ? (
            <div className="flex items-center gap-4">
              <span className="text-gray-700">
                Hello, {user.fullname || "User"}
              </span>
              <Button onClick={handleLogout} variant="secondary">
                Logout
              </Button>
            </div>
          ) : (
            <div className="flex gap-4">
              <Button onClick={() => navigate("/login")} variant="secondary">
                Login
              </Button>
              <Button onClick={() => navigate("/register")} variant="primary">
                Register
              </Button>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
