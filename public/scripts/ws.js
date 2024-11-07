export const ws = new WebSocket(`ws://${window.location.host}/ws`);

ws.onopen = () => {
  console.log("open connection");
};
