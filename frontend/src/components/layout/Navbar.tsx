import React from 'react';

const Navbar = () => {
  return (
    <nav className="bg-white border-b border-gray-200 px-6 py-4 flex justify-between items-center">
      <div className="text-xl font-semibold">TaskForge</div>
      <div className="flex items-center gap-4">
        <span>User</span>
      </div>
    </nav>
  );
};

export default Navbar;