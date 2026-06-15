import React, { useEffect, useState } from 'react';
import TaskList from '../components/tasks/TaskList';
import TaskFilters from '../components/tasks/TaskFilters';
import CreateTaskModal from '../components/tasks/CreateTaskModal';
import { fetchTasks } from '../services/api';

const Tasks = () => {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetchTasks().then(data => setTasks(data));
  }, []);

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-semibold text-gray-900">Tasks</h1>
        <CreateTaskModal />
      </div>
      <TaskFilters />
      <TaskList tasks={tasks} />
    </div>
  );
};

export default Tasks;
