const Card = ({ children, onClick, className = "" }) => {
  return (
    <div
      onClick={onClick}
      className={`bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow ${
        onClick ? "cursor-pointer" : ""
      } ${className}`}
    >
      {children}
    </div>
  );
};

export default Card;
