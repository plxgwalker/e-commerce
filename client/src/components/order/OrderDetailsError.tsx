import { Link } from "react-router-dom";

export default function OrderDetailsError() {
  return (
    <div className="flex h-screen items-center justify-center bg-gray-200 text-center">
      <div className="grid grid-cols-1">
        <p className="text-3xl font-semibold">This order is not found</p>
        <p>Make sure that provided link is correct</p>
        <p>
          Go to{" "}
          <Link to={`/`} className="underline">
            main page
          </Link>
        </p>
      </div>
    </div>
  );
}
