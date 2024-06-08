import {
  GridBackground,
  Keypad,
  MacbookScroll,
  StickyScroll,
} from "components";
import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <GridBackground>
        <StickyScroll
          content={[
            {
              title: "Chat conversation",
              description: "There are a lot of conversations",
              content: (
                <div className="h-full w-full bg-[linear-gradient(to_bottom_right,var(--orange-500),var(--yellow-500))] flex items-center justify-center text-white">
                  Chat
                </div>
              ),
            },
            {
              title: "Flows",
              description: "There are a lot of flows",
              content: (
                <div className="h-full w-full bg-[linear-gradient(to_bottom_right,var(--cyan-500),var(--emerald-500))] flex items-center justify-center text-white">
                  Flows
                </div>
              ),
            },
          ]}
        />
      </GridBackground>
    </main>
  );
}
