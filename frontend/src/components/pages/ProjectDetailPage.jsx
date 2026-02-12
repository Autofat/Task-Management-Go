import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Navbar from "../organisms/Navbar";
import TaskList from "../organisms/TaskList";
import MemberManagement from "../organisms/MemberManagement";
import Modal from "../atoms/Modal";
import TaskForm from "../molecules/TaskForm";
import Button from "../atoms/Button";
import {
  getProjectById,
  getTasks,
  createTask,
  updateTask,
  deleteTask,
  getProjectMembers,
  inviteMember,
  updateMemberRole,
  removeMember,
} from "../../services/api";

const ProjectDetailPage = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [project, setProject] = useState(null);
  const [tasks, setTasks] = useState([]);
  const [members, setMembers] = useState([]);
  const [filters, setFilters] = useState({ status: "", priority: "" });
  const [showTaskModal, setShowTaskModal] = useState(false);
  const [showMemberModal, setShowMemberModal] = useState(false);
  const [selectedTask, setSelectedTask] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  const userId = JSON.parse(localStorage.getItem("user") || "{}")?.id;
  const isOwner = project?.owner_id === userId;

  useEffect(() => {
    fetchProjectDetails();
    fetchTasks();
    fetchMembers();
  }, [id]);

  useEffect(() => {
    fetchTasks();
  }, [filters]);

  const fetchProjectDetails = async () => {
    try {
      const response = await getProjectById(id);
      setProject(response.data.data);
    } catch (error) {
      console.error("Failed to fetch project:", error);
      if (error.response?.status === 401) {
        navigate("/login");
      }
    }
  };

  const fetchTasks = async () => {
    setIsLoading(true);
    try {
      const params = { project_id: id, ...filters };
      const response = await getTasks(params);
      setTasks(response.data.data.data || []);
    } catch (error) {
      console.error("Failed to fetch tasks:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const fetchMembers = async () => {
    try {
      const response = await getProjectMembers(id);
      setMembers(response.data.data || []);
    } catch (error) {
      console.error("Failed to fetch members:", error);
    }
  };

  const handleCreateTask = async (formData) => {
    try {
      await createTask({ ...formData, project_id: parseInt(id) });
      setShowTaskModal(false);
      fetchTasks();
    } catch (error) {
      console.error("Failed to create task:", error);
    }
  };

  const handleUpdateTask = async (formData) => {
    try {
      await updateTask(selectedTask.id, formData);
      setShowTaskModal(false);
      setSelectedTask(null);
      fetchTasks();
    } catch (error) {
      console.error("Failed to update task:", error);
    }
  };

  const handleDeleteTask = async (taskId) => {
    if (window.confirm("Are you sure you want to delete this task?")) {
      try {
        await deleteTask(taskId);
        fetchTasks();
      } catch (error) {
        console.error("Failed to delete task:", error);
      }
    }
  };

  const handleInviteMember = async (userId) => {
    try {
      await inviteMember(id, parseInt(userId));
      fetchMembers();
    } catch (error) {
      console.error("Failed to invite member:", error);
    }
  };

  const handleUpdateRole = async (userId, role) => {
    try {
      await updateMemberRole(id, userId, role);
      fetchMembers();
    } catch (error) {
      console.error("Failed to update role:", error);
    }
  };

  const handleRemoveMember = async (userId) => {
    if (window.confirm("Are you sure you want to remove this member?")) {
      try {
        await removeMember(id, userId);
        fetchMembers();
      } catch (error) {
        console.error("Failed to remove member:", error);
      }
    }
  };

  const handleTaskClick = (task) => {
    setSelectedTask(task);
    setShowTaskModal(true);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="container mx-auto px-4 py-8">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
          <div>
            <Button onClick={() => navigate("/")} variant="secondary">
              ← Back to Projects
            </Button>
            <h1 className="text-3xl font-bold text-gray-800 mt-4">
              {project?.title || "Loading..."}
            </h1>
          </div>
          <Button
            onClick={() => setShowMemberModal(!showMemberModal)}
            variant="primary"
          >
            {showMemberModal ? "Hide" : "Show"} Team Members
          </Button>
        </div>

        {showMemberModal && (
          <div className="mb-8">
            <MemberManagement
              members={members}
              onInviteMember={handleInviteMember}
              onUpdateRole={handleUpdateRole}
              onRemoveMember={handleRemoveMember}
              isOwner={isOwner}
            />
          </div>
        )}

        {isLoading ? (
          <div className="text-center py-12">Loading tasks...</div>
        ) : (
          <TaskList
            tasks={tasks}
            onCreateTask={() => {
              setSelectedTask(null);
              setShowTaskModal(true);
            }}
            onTaskClick={handleTaskClick}
            onFilterChange={setFilters}
            filters={filters}
          />
        )}
      </div>

      <Modal
        isOpen={showTaskModal}
        onClose={() => {
          setShowTaskModal(false);
          setSelectedTask(null);
        }}
        title={selectedTask ? "Edit Task" : "Create New Task"}
      >
        <TaskForm
          onSubmit={selectedTask ? handleUpdateTask : handleCreateTask}
          members={members}
          initialData={selectedTask}
        />
        {selectedTask && (
          <Button
            onClick={() => handleDeleteTask(selectedTask.id)}
            variant="danger"
            className="mt-4"
          >
            Delete Task
          </Button>
        )}
      </Modal>
    </div>
  );
};

export default ProjectDetailPage;
