"use client";

import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

type Player = "X" | "O" | null;

const TicTacToe: React.FC = () => {
  const [board, setBoard] = useState<Player[]>(Array(9).fill(null));
  const [currentPlayer, setCurrentPlayer] = useState<"X" | "O">("X");
  const [winner, setWinner] = useState<Player>(null);

  const checkWinner = (squares: Player[]): Player => {
    const lines = [
      [0, 1, 2],
      [3, 4, 5],
      [6, 7, 8],
      [0, 3, 6],
      [1, 4, 7],
      [2, 5, 8],
      [0, 4, 8],
      [2, 4, 6],
    ];

    for (let i = 0; i < lines.length; i++) {
      const [a, b, c] = lines[i];
      if (
        squares[a] &&
        squares[a] === squares[b] &&
        squares[a] === squares[c]
      ) {
        return squares[a];
      }
    }

    return null;
  };

  const handleClick = (index: number) => {
    if (board[index] || winner) return;

    const newBoard = [...board];
    newBoard[index] = currentPlayer;
    setBoard(newBoard);

    const newWinner = checkWinner(newBoard);
    if (newWinner) {
      setWinner(newWinner);
    } else {
      setCurrentPlayer(currentPlayer === "X" ? "O" : "X");
    }
  };

  const resetGame = () => {
    setBoard(Array(9).fill(null));
    setCurrentPlayer("X");
    setWinner(null);
  };

  const renderSquare = (index: number) => (
    <Button
      variant="outline"
      className="w-20 h-20 text-4xl font-bold"
      onClick={() => handleClick(index)}
    >
      {board[index]}
    </Button>
  );

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle className="text-2xl font-bold text-center">
          Tic-Tac-Toe
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-3 gap-2 mb-4">
          {Array(9)
            .fill(null)
            .map((_, index) => (
              <div key={index}>{renderSquare(index)}</div>
            ))}
        </div>
        <div className="text-center mb-4">
          {winner ? (
            <p className="text-xl font-bold">Winner: {winner}</p>
          ) : board.every(Boolean) ? (
            <p className="text-xl font-bold">It's a draw!</p>
          ) : (
            <p className="text-xl">Next player: {currentPlayer}</p>
          )}
        </div>
        <div className="text-center">
          <Button onClick={resetGame}>New Game</Button>
        </div>
      </CardContent>
    </Card>
  );
};

export default TicTacToe;
