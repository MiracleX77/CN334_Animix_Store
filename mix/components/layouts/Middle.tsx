type Props = {
  children: React.ReactNode;
  X?: boolean;
  Y?: boolean;
  className?: string;
};

export default function Middle(props: Props) {
  return (
    <>
      <div
        className={`flex  ${
          props.X && props.Y
            ? "justify-center items-center"
            : props.X
            ? "justify-center items-start"
            : props.Y
            ? "justify-start items-center"
            : ""
        } ${props.className}  animate`}
      >
        {props.children}
      </div>
    </>
  );
}
