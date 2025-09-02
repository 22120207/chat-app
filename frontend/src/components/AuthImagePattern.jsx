const AuthImagePattern = ({ title, subtitle }) => {
  return (
    <div className="hidden lg:flex justify-center items-center bg-base-200">
      <div className="max-w-md text-center">
        <div className="grid grid-cols-3 gap-3 mb-5">
          {[...Array(9)].map((_, index) => (
            <div
              key={index}
              className={`aspect-square bg-primary/40 rounded-2xl ${
                index % 2 === 0 ? "animate-pulse" : ""
              }`}
            />
          ))}
        </div>

        <h2 className="text-2xl font-bold mb-4">{title}</h2>
        <p className="text-base-content/60">{subtitle}</p>
      </div>
    </div>
  );
};

export default AuthImagePattern;
