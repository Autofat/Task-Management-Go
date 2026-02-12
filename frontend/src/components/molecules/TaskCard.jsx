import Card from "../atoms/Card";
import Badge from "../atoms/Badge";

const TaskCard = ({ task, onClick }) => {
  return (
    <Card onClick={onClick}>
      <div className="flex justify-between items-start mb-2">
        <h3 className="text-lg font-semibold text-gray-800">{task.title}</h3>
        <div className="flex gap-2">
          <Badge variant={task.priority}>{task.priority}</Badge>
          <Badge variant={task.status}>{task.status.replace("_", " ")}</Badge>
        </div>
      </div>
      <p className="text-sm text-gray-600 mb-3">{task.description}</p>
      <div className="text-xs text-gray-500">
        {task.due_date && (
          <p>Due: {new Date(task.due_date).toLocaleDateString()}</p>
        )}
      </div>
    </Card>
  );
};

export default TaskCard;
