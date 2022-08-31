import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Todo } from '../../models';

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.scss'],
})
export class TodoItemComponent {
  @Input() todo: Todo | undefined;
  @Output() delete: EventEmitter<Todo> = new EventEmitter();
  @Output() markDone: EventEmitter<Todo> = new EventEmitter();

  deleteTodo(): void {
    this.delete.emit(this.todo);
  }

  markCompleted(): void {
    this.markDone.emit(this.todo);
  }
}
