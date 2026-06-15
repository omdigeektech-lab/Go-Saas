import React from 'react';

const TaskList = ({ tasks }: { tasks: any[] }) => {
  return (
    <div className="bg-white shadow rounded-lg p-4">
      {tasks.length === 0 ? (
        <p>No tasks found.</p>
      ) : (
        <ul>
          {tasks.map((task, idx) => (
            <li key={idx} className="border-b py-2">{task.title}</li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default TaskList;
