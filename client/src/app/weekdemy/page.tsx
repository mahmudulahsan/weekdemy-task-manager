'use client';
import { useEffect, useState } from 'react';
import Navbar from '../components/Navbar';

export default function Page() {
  const [tasks, setTasks] = useState([]);
  const [newTeam, setNewTeam] = useState({
    teamName: '',
    projectName: '',
    isFinished: false,
    startTime: '',
    finishedTime: '',
  });
  const [editingTaskId, setEditingTaskId] = useState(null);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await fetch('http://192.168.11.150:8080/weekdemy/teams');
        const data = await response.json();
        setTasks(data);
      } catch (error) {
        console.error('Error fetching tasks:', error);
      }
    };
    fetchTasks();
  }, []);

  const handleInputChange = (e) => {
    setNewTeam((prevTeam) => ({
      ...prevTeam,
      [e.target.name]: e.target.value,
    }));
  };

  const handleAddTeam = async (e) => {
    e.preventDefault();
    try {
      // Insert the current time in the startTime field
      const currentTime = new Date();
      const formattedTime = currentTime.toISOString();
  
      const response = await fetch('http://192.168.11.150:8080/weekdemy/teams', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ...newTeam,
          startTime: formattedTime,
        }),
      });
  
      if (response.ok) {
        // Successfully created team, fetch updated tasks
        const updatedResponse = await fetch('http://192.168.11.150:8080/weekdemy/teams');
        const updatedData = await updatedResponse.json();
        setTasks(updatedData);
        // Clear the form after successful addition
        setNewTeam({
          teamName: '',
          projectName: '',
          isFinished: false,
          startTime: '',
          finishedTime: '',
        });
        console.log('Team added successfully');
      } else {
        console.error('Failed to add team');
      }
    } catch (error) {
      console.error('Error adding team:', error);
    }
  };

  const handleDeleteTeam = async (taskId) => {
    try {
      const response = await fetch(`http://192.168.11.150:8080/weekdemy/teams/${taskId}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        // Successfully deleted team, fetch updated tasks
        const updatedResponse = await fetch('http://192.168.11.150:8080/weekdemy/teams');
        const updatedData = await updatedResponse.json();
        setTasks(updatedData);
        console.log('Team deleted successfully');
      } else {
        console.error('Failed to delete team');
      }
    } catch (error) {
      console.error('Error deleting team:', error);
    }
  };

  const handleUpdateTeam = async () => {
    try {
      const response = await fetch(`http://192.168.11.150:8080/weekdemy/teams/${editingTaskId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newTeam),
      });

      if (response.ok) {
        // Successfully updated team, fetch updated tasks
        const updatedResponse = await fetch('http://192.168.11.150:8080/weekdemy/teams');
        const updatedData = await updatedResponse.json();
        setTasks(updatedData);
        // Clear the form and reset editingTaskId after successful update
        setNewTeam({
          teamName: '',
          projectName: '',
          isFinished: false,
          startTime: '',
          finishedTime: '',
        });
        setEditingTaskId(null);
        console.log('Team updated successfully');
      } else {
        console.error('Failed to update team');
      }
    } catch (error) {
      console.error('Error updating team:', error);
    }
  };

  const startEditingTeam = (taskId) => {
    // Find the task by ID and set its details in the form for editing
    const taskToEdit = tasks.find((task) => task.id === taskId);
    setNewTeam({
      teamName: taskToEdit.teamName,
      projectName: taskToEdit.projectName,
      isFinished: taskToEdit.isFinished,
      startTime: taskToEdit.startTime,
      finishedTime: taskToEdit.finishedTime,
    });
    setEditingTaskId(taskId);
  };

  return (
    <main>
      <Navbar />
      <h1 className="mt-10 text-5xl font-bold text-center">Weekdemy Task Manager</h1>

      <div className='mt-10 flex justify-center items-center'>
      {/* Form for adding/updating a team */}
      <div className="p-10 bg-base-300 flex flex-col items-center mt-4">
        {editingTaskId ? (<h1 className='font-bold text-2xl text-info'>Update the Team</h1>):(<h1 className="font-bold text-2xl text-info">Add The Team</h1>)}
      <form>
          <label>
            <input type="text" name="teamName"
              value={newTeam.teamName}
              onChange={handleInputChange} 
              placeholder="Team Name" 
              className="m-3 input input-bordered w-full max-w-xs" />
          </label> <br />
          <label>
            <input type="text" name="projectName"
              value={newTeam.projectName}
              onChange={handleInputChange} 
              placeholder="Project Name" 
              className="m-3 input input-bordered w-full max-w-xs" />
          </label>
          <div className="m-4 flex justify-center">
        {/* <button className="btn btn-secondary" onClick={handleAddTeam}>
          Add Team
        </button> */}
        <button className="btn btn-success" onClick={editingTaskId ? handleUpdateTeam : handleAddTeam}>
          {editingTaskId ? 'Update Team' : 'Add Team'}
        </button>
      </div>
        </form>
      </div>

      {/* Display the fetched tasks */}
      <div className="flex justify-center items-center">
        <div className="overflow-x-auto">
          <table className="table">
            <thead>
              <tr>
                <th></th>
                <th>Team Name</th>
                <th>Project Name</th>
                <th>Completed?</th>
                <th>Start Time</th>
                <th>Finished Time</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {tasks.map((task) => (
                <tr key={task.id}>
                  <th>{task.id}</th>
                  <td>{task.teamName}</td>
                  <td>{task.projectName}</td>
                  <td>
                    <button
                      className={`btn ${task.isFinished ? 'btn-error' : 'btn-success'}`}
                    >
                      {task.isFinished ? 'Mark as undone' : 'Mark as done'}
                    </button>
                  </td>
                  <td>{task.startTime}</td>
                  <td>{task.finishedTime}</td>
                  <td>
                    <button
                      className="m-2 btn btn-xs btn-primary"
                      onClick={() => startEditingTeam(task.id)}
                    >
                      Update
                    </button>
                    <button
                      className="m-2 btn btn-xs btn-error"
                      onClick={() => handleDeleteTeam(task.id)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
      </div>
    </main>
  );
}
