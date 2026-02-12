import TaskCard from "../molecules/TaskCard";
import Button from "../atoms/Button";
import Select from "../atoms/Select";

const TaskList = ({
  tasks,
  onCreateTask,
  onTaskClick,
  onFilterChange,
  filters,
}) => {
  const statusOptions = [
    { value: "", label: "All Status" },
    { value: "pending", label: "Pending" },
    { value: "in_progress", label: "In Progress" },
    { value: "completed", label: "Completed" },
  ];

  const priorityOptions = [
    { value: "", label: "All Priority" },
    { value: "low", label: "Low" },
    { value: "medium", label: "Medium" },
    { value: "high", label: "High" },
  ];

  return (
    <div>
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Tasks</h2>
        <div className="flex flex-col md:flex-row gap-4 w-full md:w-auto">
          <Select
            name="status"
            value={filters?.status || ""}
            onChange={(e) =>
              onFilterChange({ ...filters, status: e.target.value })
            }
            options={statusOptions}
          />
          <Select
            name="priority"
            value={filters?.priority || ""}
            onChange={(e) =>
              onFilterChange({ ...filters, priority: e.target.value })
            }
            options={priorityOptions}
          />
          <Button onClick={onCreateTask} variant="primary">
            + New Task
          </Button>
        </div>
      </div>

      {tasks.length === 0 ? (
        <div className="text-center py-12 text-gray-500">
          <p>No tasks found. Create your first task!</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {tasks.map((task) => (
            <TaskCard
              key={task.id}
              task={task}
              onClick={() => onTaskClick(task)}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default TaskList;
