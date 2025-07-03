import { Container, getContainer } from "@cloudflare/containers";

export class DebtRecyclingCalc extends Container {
  defaultPort = 8080;
  sleepAfter = "2m";
  envVars = {};

  onStart() {
    console.log("container successfully started");
  }

  onStop() {
    console.log("container successfully shut down");
  }

  onError(error) {
    console.error("error:", error);
  }
}

export default {
  async fetch(request, env, ctx) {
    try {
      const sessionId = "default-session";

      const containerInstance = getContainer(env.MY_CONTAINER, sessionId);

      return containerInstance.fetch(request);
    } catch (error) {
      console.error("Error in worker fetch handler:", error);

      return new Response(
        `Worker Error: ${error instanceof Error ? error.message : String(error)}`,
        { status: 500 },
      );
    }
  },
};
