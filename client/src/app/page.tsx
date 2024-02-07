import Link from "next/link";
import Navbar from "./components/Navbar";


export default function Home() {
  return (
    <main className="">
      <Navbar/>
      <div className="flex justify-center">
        <Link href="/weekdemy" className="mt-16 btn btn-lg uppercase btn-warning">goto weekdemy Dashboard</Link>
       </div>
    </main>
  );
}
