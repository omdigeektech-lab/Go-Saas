import React from 'react';

const Badge = ({ children }: any) => {
  return (
    <span className="bg-blue-100 text-blue-800 text-xs font-semibold px-2.5 py-0.5 rounded">
      {children}
    </span>
  );
};

export default Badge;