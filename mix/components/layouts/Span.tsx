type Props = {
  children: React.ReactNode;
  X?: boolean;
  Y?: boolean;
  className?: string;
};

export default function Span(props: Props) {
  return (
    <>
      <span
        className={`flex ${
          props.X && props.Y
            ? "justify-between items-center"
            : props.X
            ? "justify-around items-start"
            : props.Y
            ? "flex-col justify-around items-start"
            : ""
        } ${props.className} animate`}
      >
        {props.children}
      </span>
    </>
  );
}