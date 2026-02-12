import { useNavigate } from "react-router-dom";
import ProjectCard from "../molecules/ProjectCard";
import Button from "../atoms/Button";

const ProjectList = ({ projects, onCreateProject }) => {
  const navigate = useNavigate();

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">My Projects</h2>
        <Button onClick={onCreateProject} variant="primary">
          + New Project
        </Button>
      </div>

      {projects.length === 0 ? (
        <div className="text-center py-12 text-gray-500">
          <p>No projects found. Create your first project!</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {projects.map((project) => (
            <ProjectCard
              key={project.id}
              project={project}
              onClick={() => navigate(`/projects/${project.id}`)}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default ProjectList;
