import React from 'react';

const TaskFilters = () => {
  return (
    <div className="flex gap-4 mb-4">
      <select className="border rounded p-2">
        <option>All Status</option>
        <option>Pending</option>
        <option>Completed</option>
      </select>
    </div>
  );
};

export default TaskFilters;
