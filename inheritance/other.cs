namespace Foo
{
    public abstract class Foo
    {
        public int FooBar { get; set; }
    }

    public class Bar : Foo
    {
        public int Baz()
        {
            return base.FooBar;
        }
    }
}