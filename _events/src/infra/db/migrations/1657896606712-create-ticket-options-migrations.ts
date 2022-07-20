import { MigrationInterface, QueryRunner, Table, TableForeignKey } from 'typeorm'

export class CreateTicketOptionsMigration implements MigrationInterface {
  name = this.constructor.name.concat('1657896606712')

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'ticket_options',
        columns: [
          { name: 'id', type: 'uuid', isPrimary: true },
          { name: 'event_id', type: 'uuid', isNullable: false },
          { name: 'title', type: 'varchar', isNullable: false },
          { name: 'price', type: 'decimal', precision: 10, scale: 2, isNullable: false },
          { name: 'description', type: 'text', isNullable: true },
          { name: 'maximum_buys_quantity', type: 'int', isNullable: false },
          { name: 'minimum_buys_quantity', type: 'int', isNullable: false },
          { name: 'total_quantity', type: 'int', isNullable: false },
          { name: 'remaining_quantity', type: 'int', isNullable: false },
          { name: 'created_at', type: 'timestamp', default: 'now()' },
          { name: 'updated_at', type: 'timestamp', default: 'now()' }
        ]
      })
    )

    await queryRunner.createForeignKey(
      'ticket_options',
      new TableForeignKey({
        name: 'fk_events_ticket_options',
        referencedTableName: 'events',
        referencedColumnNames: ['id'],
        columnNames: ['event_id'],
        onDelete: 'CASCADE',
        onUpdate: 'CASCADE'
      })
    )
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropForeignKey('ticket_options', 'fk_events_ticket_options')
    await queryRunner.dropTable('ticket_options', true)
  }
}
