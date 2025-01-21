// src/components/Stopwatch.tsx
import { useState, useEffect } from 'react';

export default function Stopwatch() {
  const [time, setTime] = useState(0);
  const [isRunning, setIsRunning] = useState(false);
  const [laps, setLaps] = useState<number[]>([]);

  useEffect(() => {
    let intervalId: number;
    if (isRunning) {
      intervalId = window.setInterval(() => {
        setTime((prevTime) => prevTime + 10);
      }, 10);
    }
    return () => clearInterval(intervalId);
  }, [isRunning]);

  const handleStartStop = () => {
    setIsRunning(!isRunning);
  };

  const handleReset = () => {
    setIsRunning(false);
    setTime(0);
    setLaps([]);
  };

  const handleLap = () => {
    setLaps([...laps, time]);
  };

  const formatTime = (ms: number) => {
    const minutes = Math.floor(ms / 60000);
    const seconds = Math.floor((ms % 60000) / 1000);
    const milliseconds = Math.floor((ms % 1000) / 10);

    return `${minutes.toString().padStart(2, '0')}:${seconds
      .toString()
      .padStart(2, '0')}.${milliseconds.toString().padStart(2, '0')}`;
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-lg shadow-lg">
      <div className="text-6xl font-mono text-center mb-8">
        {formatTime(time)}
      </div>
      
      <div className="flex justify-center gap-4 mb-6">
        <button
          onClick={handleStartStop}
          className={`px-6 py-2 rounded-lg ${
            isRunning
              ? 'bg-red-500 hover:bg-red-600'
              : 'bg-green-500 hover:bg-green-600'
          } text-white`}
        >
          {isRunning ? 'Stop' : 'Start'}
        </button>
        
        <button
          onClick={handleLap}
          disabled={!isRunning}
          className="px-6 py-2 rounded-lg bg-blue-500 hover:bg-blue-600 text-white disabled:opacity-50"
        >
          Lap
        </button>
        
        <button
          onClick={handleReset}
          className="px-6 py-2 rounded-lg bg-gray-500 hover:bg-gray-600 text-white"
        >
          Reset
        </button>
      </div>

      {laps.length > 0 && (
        <div className="mt-6">
          <h3 className="text-lg font-semibold mb-2">Laps</h3>
          <div className="space-y-2">
            {laps.map((lapTime, index) => (
              <div
                key={index}
                className="flex justify-between bg-gray-100 p-2 rounded"
              >
                <span>Lap {index + 1}</span>
                <span>{formatTime(lapTime)}</span>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}