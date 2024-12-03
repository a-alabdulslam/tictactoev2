import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";

const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  gamecode: z.string().length(6, {
    message: "Game code must be exactly 6 characters.",
  }),
});

export function Homepage() {
  // 1. Define your form.
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      gamecode: "",
    },
  });
  const navigate = useNavigate();
  // 2. Define a submit handler.
  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    navigate("/game");
    console.log(values);
  }
  // Submit handler for creating a new game.
  function createGame() {
    const { username } = form.getValues(); // Only retrieve the username.
    if (!username || username.length < 2) {
      form.setError("username", {
        type: "manual",
        message: "Username must be at least 2 characters.",
      });
      return;
    }
    navigate("/game");

    console.log("Create Game with Username:", username);
  }
  return (
    <>
      <h1 className="text-2xl font-bold">Join a game</h1>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder="Abdul" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="gamecode"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Gamecode</FormLabel>
                <FormControl>
                  <Input placeholder="######" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">Join</Button>{" "}
          <Button type="button" onClick={createGame}>
            Create Game
          </Button>
        </form>
      </Form>
    </>
  );
}
