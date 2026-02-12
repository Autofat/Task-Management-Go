import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../organisms/Navbar";
import ProjectList from "../organisms/ProjectList";
import Modal from "../atoms/Modal";
import ProjectForm from "../molecules/ProjectForm";
import { getProjects, createProject } from "../../services/api";

const HomePage = () => {
  const navigate = useNavigate();
  const [projects, setProjects] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [isCreating, setIsCreating] = useState(false);

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    setIsLoading(true);
    try {
      const response = await getProjects();
      setProjects(response.data.data.data || []);
    } catch (error) {
      console.error("Failed to fetch projects:", error);
      if (error.response?.status === 401) {
        navigate("/login");
      }
    } finally {
      setIsLoading(false);
    }
  };

  const handleCreateProject = async (formData) => {
    setIsCreating(true);
    try {
      await createProject(formData.title);
      setShowModal(false);
      fetchProjects();
    } catch (error) {
      console.error("Failed to create project:", error);
    } finally {
      setIsCreating(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="container mx-auto px-4 py-8">
        {isLoading ? (
          <div className="text-center py-12">Loading...</div>
        ) : (
          <ProjectList
            projects={projects}
            onCreateProject={() => setShowModal(true)}
          />
        )}
      </div>

      <Modal
        isOpen={showModal}
        onClose={() => setShowModal(false)}
        title="Create New Project"
      >
        <ProjectForm onSubmit={handleCreateProject} isLoading={isCreating} />
      </Modal>
    </div>
  );
};

export default HomePage;
