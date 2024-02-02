

import Link from "next/link";
import Navbar from "./components/Navbar"


export default function Home() {
  return (
    <main className="">
      <Navbar/>
      This is home <br />
      <Link href="/weekdemy" className="btn btn-lg">goto weekdemy route</Link>
    </main>
  );
}
