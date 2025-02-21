import { useEffect, useState } from "react";

export default function Dashboard() {
	const [tasks, setTasks] = useState([]);

	useEffect(() => {
		const ws = new WebSocket("ws://localhost:8080/ws");
		ws.onmessage = (event) => {
			setTasks((prevTasks) => [...prevTasks, JSON.parse(event.data)]);
		};
	}, []);

	return (
		<div>
			<h1>Task Dashboard</h1>
			<ul>
				{tasks.map((task) => (
					<li key={task.id}>{task.title}</li>
				))}
			</ul>
		</div>
	);
}
