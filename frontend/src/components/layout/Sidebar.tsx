import React from 'react';
import { Link } from 'react-router-dom';

const Sidebar = () => {
  return (
    <div className="w-64 bg-gray-900 text-white flex flex-col">
      <div className="p-6 text-2xl font-bold">Go-Saas</div>
      <nav className="flex-1 px-4 py-6 space-y-2">
        <Link to="/" className="block px-4 py-2 rounded hover:bg-gray-800">Dashboard</Link>
        <Link to="/tasks" className="block px-4 py-2 rounded hover:bg-gray-800">Tasks</Link>
      </nav>
    </div>
  );
};

export default Sidebar;