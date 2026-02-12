import Card from "../atoms/Card";

const ProjectCard = ({ project, onClick }) => {
  return (
    <Card onClick={onClick}>
      <h3 className="text-xl font-semibold text-gray-800 mb-2">
        {project.title}
      </h3>
      <div className="text-sm text-gray-600">
        <p>Owner: {project.owner?.fullname || "Unknown"}</p>
        <p className="text-xs text-gray-400 mt-2">
          Created: {new Date(project.created_at).toLocaleDateString()}
        </p>
      </div>
    </Card>
  );
};

export default ProjectCard;
