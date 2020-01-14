package main

import "testing"

func BenchmarkRender2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Render(2)
    }
}

func BenchmarkRender4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Render(4)
    }
}

func BenchmarkRender6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Render(6)
    }
}

func BenchmarkRender8(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Render(8)
    }
}

func BenchmarkRender10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Render(10)
    }
}
