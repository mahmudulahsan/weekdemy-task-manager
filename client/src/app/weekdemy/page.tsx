// import Navbar from "../components/Navbar"

// export default function page() {
//     return (
//       <main className="">
//         <Navbar/>
//         <h1 className="text-6xl font-bold text-center">Weekdemy Task Manager</h1>
//       </main>
//     );
//   }

'use client';
import { useEffect, useState } from "react";
import Navbar from "../components/Navbar";

export default function Page() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await fetch("http://192.168.11.150:8080/weekdemy/teams");
        const data = await response.json();
        setTasks(data);
        // console.log("Fetched tasks:", data)
      } catch (error) {
        console.error("Error fetching tasks:", error);
      }
    };
    fetchTasks();
  }, []);

  return (
    <main className="">
      <Navbar />
      <h1 className="mt-10 text-5xl font-bold text-center">Weekdemy Task Manager</h1>

      <div className="overflow-x-auto">
  <table className="table">
    {/* head */}
    <thead>
      <tr>
        <th></th>
        <th>Team Name</th>
        <th>Project Name</th>
        <th>Completed?</th>
        <th>Start Time</th>
        <th>End Time</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
    {tasks.map((task) => (
          <tr key={task.id}>
            <th>{task.id}</th>
            <td>{task.teamName}</td>
            <td>{task.projectName}</td>
            <td>{task.isFinished ? "Yes" : "No"}</td>
            <td>-</td>
            <td>-</td>
            {/* {task.isFinished && <p>Finished Time: {task.finishedTime}</p>} */}
            <button className="btn btn-primary">update</button>
            <button className="btn btn-error">delete</button>
          </tr>
        ))}
      {/* <tr className="bg-base-200">
        <th>1</th>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
      </tr>

      <tr>
        <th>2</th>
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
      </tr>

      <tr>
        <th>3</th>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
      </tr> */}
    </tbody>
  </table>
</div>


      
        {/* {tasks.map((task) => (
          <tr key={task.id}>
            <th>1</th>
            <td>Team: {task.teamName}</td>
            <td>Project: {task.projectName}</td>
            <td>Finished: {task.isFinished ? "Yes" : "No"}</td>
            {task.isFinished && <p>Finished Time: {task.finishedTime}</p>}
          </tr>
        ))} */}
    </main>
  );
}
